package program

import (
	"log"

	"github.com/koodinikkarit/matias/configuration"
	"github.com/koodinikkarit/matias/generators"
	"github.com/koodinikkarit/matias/matiasdatabase"
)

func (p *Program) parseConfig() {
	configFile := configuration.ConfigFile{}
	if fileExits("config.yml") == true {
		log.Println("config.yml configfile found")
		configFile.Load("config.yml")
	} else {
		log.Println("config.yml not found")
	}

	if configFile.MatiasKey == "" {
		randomString, err := generators.GenerateRandomString(10)
		if err != nil {
			log.Fatalf("Error when creating key %v", err)
		}
		configFile.MatiasKey = randomString
		p.matiasKey = configFile.MatiasKey
	}

	if configFile.MatiasDatabasePath == "" {
		configFile.MatiasDatabasePath = "matiasdatabase.db"
	}

	if fileExits(configFile.MatiasDatabasePath) == false {
		log.Printf("%v ei loydy luodaan matiasdatabase.db", configFile.MatiasDatabasePath)
		configFile.MatiasDatabasePath = "matiasdatabase.db"
	}

	if p.matiasDatabase == nil {
		p.matiasDatabase = matiasdatabase.NewMatiasDatabase(
			configFile.MatiasDatabasePath,
		)
	} else if p.configFile.MatiasDatabasePath != configFile.MatiasDatabasePath {
		p.matiasDatabase.Close()
		p.matiasDatabase = matiasdatabase.NewMatiasDatabase(
			configFile.MatiasDatabasePath,
		)
	}
	log.Println(configFile)
	if configFile.MatiasServiceIP == "" {
		log.Printf("Configfile does not contain matiasServiceIp connection to server cannot be created")
	}

	if configFile.MataisServicePort == "" {
		log.Printf("Configfile does not contain matiasServicePort using default 6755 instead")
		configFile.MataisServicePort = "6755"
	}

	if configFile.MatiasServiceIP != p.configFile.MatiasServiceIP ||
		configFile.MataisServicePort != p.configFile.MataisServicePort {
		p.matiasServicePort = configFile.MataisServicePort
		if p.matiasClient != nil {
			p.matiasClient.Close()
			p.matiasClient = nil
		}
		if configFile.MatiasServiceIP != "" {
			p.matiasServiceIP = configFile.MatiasServiceIP
			p.createMatiasClient()
		}
	}

	ewDatabases := p.matiasDatabase.GetEwDatabases()
	for _, ewDatabase := range ewDatabases {
		p.GetOrCreateEwDatabaseInstance(
			ewDatabase,
		)
	}

	// // ssl ??
	// if configFile.Ssl == true {

	// }

	log.Println("Tallentetaan config.yml")
	configFile.Save("config.yml")
	p.configFile = configFile
}
