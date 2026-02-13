package config

func (c *Config) Validate() error {

	if err := required(c.AppPort, "APP_PORT"); err != nil {
		return err
	}

	if err := required(c.ServiceName, "SERVICE_NAME"); err != nil {
		return err
	}

	if err := required(c.ServiceVersion, "SERVICE_VERSION"); err != nil {
		return err
	}

	return nil
}
