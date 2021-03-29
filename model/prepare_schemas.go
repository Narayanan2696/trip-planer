package model

func PrepareSchemas() {
	createAutomobileDetails()
	createLocations()
}

func createAutomobileDetails() {
	automobileDetailsSchema := "CREATE TABLE IF NOT EXISTS automobileDetails (id SERIAL PRIMARY KEY, mpg varchar(50), cylinders varchar(50), displacement varchar(50), horse_power varchar(50), weight varchar(50), acceleration varchar(50), model_year varchar(50), origin varchar(50), car_name varchar(450), fuel_type varchar(50))"
	_, err := connect.Exec(automobileDetailsSchema)
	if err != nil {
		panic(err)
	}
}

func createLocations() {
	locationsSchema := "CREATE TABLE IF NOT EXISTS locations (id SERIAL PRIMARY KEY, place varchar(450), latitude decimal(9,6), longitude decimal(9,6))"
	_, err := connect.Exec(locationsSchema)
	if err != nil {
		panic(err)
	}
}