package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var clusters Clusters

func init() {
	clusters = loadJson()
}

type Clusters struct {
	Clusters []Cluster `json:"clusters"`
}

type Cluster struct {
	Id             string `json:"id"`
	Description    string `json:"description"`
	Url            string `json:"cluster_url"`
	User           string `json:"user"`
	Password       string `json:"password"`
	DefaultProject string `json:"default_project"`
	CurrentContext bool   `json:"current_context"`
}

func (cluster Cluster) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Description: %s \n", cluster.Id))
	sb.WriteString(fmt.Sprintf("Description: %s \n", cluster.Description))
	sb.WriteString(fmt.Sprintf("Cluster URL: %s \n", cluster.Url))
	sb.WriteString(fmt.Sprintf("Password: %s \n", "******"))
	sb.WriteString(fmt.Sprintf("Default Project: %s \n", cluster.DefaultProject))
	sb.WriteString(fmt.Sprintf("Current Context: %t \n", cluster.CurrentContext))

	return sb.String()
}

func (clusters Clusters) String() string {
	encryptPassword(clusters)
	json, _ := json.MarshalIndent(clusters, "", "    ")
	return string(json)
}

func loadJson() Clusters {
	var clusters Clusters
	jsonFile, err := os.Open(getConfigPath())
	if err != nil {
		fmt.Println("Failed to load config.json under ~/.oclogin/ folder, please init with the command oclogin init and update the file based on your config ")
	} else {
		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(byteValue, &clusters)
		if err != nil {
			panic(err)
		}
		defer jsonFile.Close()
	}
	return clusters
}

func UpdateContext(id string) {
	var clusterArray []Cluster
	anyIdMatch := false
	for _, cluster := range clusters.Clusters {
		if cluster.Id == id {
			cluster.CurrentContext = true
			anyIdMatch = true
		} else {
			cluster.CurrentContext = false
		}
		clusterArray = append(clusterArray, cluster)
	}

	if !anyIdMatch {
		fmt.Println("Failed to find any matched id " + id)
		return
	}

	clusters = Clusters{Clusters: clusterArray}
	byteArray, err := json.Marshal(clusters)
	if err != nil {
		panic(err)
	}
	outputPath := getConfigPath()
	ioutil.WriteFile(outputPath, formatOutput(byteArray), 0755)
	fmt.Println(clusters.String())
	clusters = loadJson()
}


func Login() {
	for _, cluster := range clusters.Clusters {
		if cluster.CurrentContext {
			response, err := exec.Command("oc", "login", cluster.Url, "-u", cluster.User, "-p", cluster.Password).CombinedOutput()
			if err != nil {
				fmt.Println(string(response))
				panic(err)
			}
			fmt.Println(string(response))

			response, err = exec.Command("oc", "project", cluster.DefaultProject).CombinedOutput()
			if err != nil {
				fmt.Println(string(response))
				panic(err)
			}
			fmt.Print(string(response))
		}
	}
}

func List() {
	var clusterArray []Cluster
	for _, cluster := range clusters.Clusters {
		cluster.Password = "*******"
		clusterArray = append(clusterArray, cluster)
	}

	tmpClusters := Clusters{Clusters: clusterArray}
	fmt.Println(tmpClusters.String())
}

func Init() {
	if _, err := os.Stat(getConfigPath()); os.IsNotExist(err) {
		fmt.Println("Config file is not existed, sample will be created at the ~/.oclogin/config.json")
		templateConfig, readErr := os.ReadFile("config.json")
		if readErr != nil {
			panic(readErr)
		}
		ensureDir(getConfigPath())
		err = ioutil.WriteFile(getConfigPath(), formatOutput(templateConfig), 0755)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("Config file is already existed, skip initialization")
	}
}

func getConfigPath() string{
	homeDir, _ := os.UserHomeDir()
	return homeDir + string(os.PathSeparator) + ".oclogin" + string(os.PathSeparator) + "config.json"
}

func encryptPassword(clusters Clusters) {
	for _, cluster := range clusters.Clusters {
		cluster.Password = "******"
	}
}

func ensureDir(fileName string) {
	dirName := filepath.Dir(fileName)
	if _, serr := os.Stat(dirName); serr != nil {
		merr := os.MkdirAll(dirName, os.ModePerm)
		if merr != nil {
			panic(merr)
		}
	}
}

func formatOutput(rawData []byte) []byte {
	var tmpObj Clusters
	ensureDir(getConfigPath())
	json.Unmarshal(rawData, &tmpObj)
	formattedOutput, _ := json.MarshalIndent(tmpObj, "", "    ")
	return formattedOutput
}
