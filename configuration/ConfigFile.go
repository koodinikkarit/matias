package configuration

type ConfigFile struct {
	MatiasKey          string `yaml:"matiasKey"`
	MatiasDatabasePath string `yaml:"matiasDatabasePath"`
	MatiasServiceIP    string `yaml:"matiasServiceIp"`
	MataisServicePort  string `yaml:"matiasServicePort"`
	PrivateKeyPath     string `yaml:"privateKeyPath"`
	PublicKeyPath      string `yaml:"publicKeyPath"`
	Ssl                bool   `yaml:"ssl"`
}
