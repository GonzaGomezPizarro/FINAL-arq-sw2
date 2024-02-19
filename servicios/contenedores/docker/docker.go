package docker

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/contenedores/dto"
)

func GetContenedores() (dto.Contenedores, error) {
	// Ejecutar el comando 'docker ps -a' para obtener la lista de todos los contenedores
	cmd := exec.Command("docker", "ps", "-a", "--format", "{{.ID}}|{{.Names}}|{{.Image}}|{{.Status}}")
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar el comando docker ps -a: %v", err)
	}

	// Analizar la salida del comando para obtener la información de los contenedores
	var contenedores []dto.Contenedor
	lines := strings.Split(strings.TrimSpace(string(out)), "\n") // Eliminar espacios en blanco adicionales y dividir por líneas
	for _, line := range lines {
		if line != "" {
			fields := strings.Split(line, "|")
			contenedor := dto.Contenedor{
				Id:           fields[0],
				Name:         fields[1],
				Imagen:       dto.Imagen{Name: fields[2]},
				Status:       fields[3],
				InternalPort: 0, // Asignar un valor predeterminado de puerto interno
				ExternalPort: 0, // Asignar un valor predeterminado de puerto externo
			}
			contenedores = append(contenedores, contenedor)
		}
	}

	return contenedores, nil
}

func GetImages() (dto.Imagenes, error) {
	// Ejecutar el comando 'docker images' para obtener la lista de imágenes
	cmd := exec.Command("docker", "images")
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar el comando docker images: %v", err)
	}

	// Dividir la salida en líneas
	lines := strings.Split(string(out), "\n")

	// Construir la lista de estructuras Imagen
	var imagenes dto.Imagenes
	for _, line := range lines[1:] { // Omitir la primera línea que contiene encabezados
		fields := strings.Fields(line)
		if len(fields) >= 1 {
			imagenes = append(imagenes, dto.Imagen{Name: fields[0]})
		}
	}

	return imagenes, nil
}

func PostContenedor(contenedor dto.Contenedor) (dto.Contenedor, error) {
	// Construir los argumentos para el comando docker run
	args := []string{"run", "-d", "--network=servicios_mi_red"}
	if contenedor.Name != "" {
		args = append(args, "--name", contenedor.Name)
	}
	if contenedor.ExternalPort != 0 && contenedor.InternalPort != 0 {
		args = append(args, "-p", fmt.Sprintf("%d:%d", contenedor.ExternalPort, contenedor.InternalPort))
	}
	args = append(args, contenedor.Imagen.Name)

	// Ejecutar el comando docker run
	cmd := exec.Command("docker", args...)
	out, err := cmd.Output()
	if err != nil {
		return dto.Contenedor{}, fmt.Errorf("error al ejecutar el comando docker run: %v", err)
	}

	// Obtener el ID del contenedor creado
	containerID := strings.TrimSpace(string(out))

	// Intentar obtener el nombre del contenedor hasta que sea exitoso o se alcance el límite de intentos
	var containerName string
	attempts := 0
	maxAttempts := 10
	for containerName == "" && attempts < maxAttempts {
		attempts++

		// Obtener el nombre del contenedor creado
		containerNameCmd := exec.Command("docker", "inspect", "--format='{{.Name}}'", containerID)
		containerNameOut, err := containerNameCmd.Output()
		if err == nil {
			containerName = strings.Trim(string(containerNameOut), "'/\n")
		} else {
			time.Sleep(250 * time.Millisecond) // Esperar un corto tiempo antes de intentar nuevamente
		}
	}

	if containerName == "" {
		return dto.Contenedor{}, fmt.Errorf("no se pudo obtener el nombre del contenedor después de %d intentos", maxAttempts)
	}

	// Obtener información del contenedor creado
	containerInfoCmd := exec.Command("docker", "inspect", containerID)
	containerInfoOut, err := containerInfoCmd.Output()
	if err != nil {
		return dto.Contenedor{}, fmt.Errorf("error al obtener información del contenedor: %v", err)
	}

	// Analizar la salida del comando 'docker inspect' para obtener los detalles del contenedor
	var containerInfo []map[string]interface{}
	if err := json.Unmarshal(containerInfoOut, &containerInfo); err != nil {
		return dto.Contenedor{}, fmt.Errorf("error al analizar la información del contenedor: %v", err)
	}

	// Extraer la información necesaria del contenedor
	contenedor.Name = containerName
	contenedor.Id = containerID
	contenedor.Status = containerInfo[0]["State"].(map[string]interface{})["Status"].(string)

	return contenedor, nil
}

func StartContenedor(id string) error {
	// Ejecutar el comando 'docker start' para iniciar el contenedor con el ID proporcionado
	cmd := exec.Command("docker", "start", id)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error al ejecutar el comando docker start: %v", err)
	}
	return nil
}

func StopContenedor(id string) error {
	// Ejecutar el comando 'docker stop' para detener el contenedor con el ID proporcionado
	cmd := exec.Command("docker", "stop", id)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error al ejecutar el comando docker stop: %v", err)
	}
	return nil
}

func DeleteContenedor(id string) error {
	// Detener el contenedor antes de intentar eliminarlo
	if err := StopContenedor(id); err != nil {
		return fmt.Errorf("error al detener el contenedor: %v", err)
	}

	// Ejecutar el comando 'docker rm' para eliminar el contenedor con el ID proporcionado
	cmd := exec.Command("docker", "rm", id)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error al ejecutar el comando docker rm: %v", err)
	}
	return nil
}
