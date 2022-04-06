# ======================================================= #
# This code snippet shows the functionality of Docstring. #
# ======================================================= #

class Sensor_v1():

    def __init__(self):
        self.temp = 0.0
        self.pres = 0.0
        self.rh   = 0.0

    def measure(self, temp, pres, rh):
        self.temp = temp
        self.pres = pres
        self.rh   = rh
    
    def display_temperature(self):
        print(f"Temperature: {self.temp}")

    def display_pressure(self):
        print(f"Pressure: {self.pres}")

    def display_relative_humidity(self):
        print(f"Relative Humidity: {self.rh}")

class Sensor_v2():

    """
    This class contains methods to detect environmental variables.
    """

    def __init__(self):
        """
        Provide default values to target variables.
        """
        self.temp = 0.0
        self.pres = 0.0
        self.rh   = 0.0

    def measure(self, temp, pres, rh):
        """
        Sensor will open up and measure temperature, pressure, 
        and relative humidity
        """
        self.temp = temp
        self.pres = pres
        self.rh   = rh
    
    def display_temperature(self):
        """
        Display detected temperature.
        """
        print(f"Temperature: {self.temp}")

    def display_pressure(self):
        """
        Display detected pressure.
        """
        print(f"Pressure: {self.pres}")

    def display_relative_humidity(self):
        """
        Display relative humidity.
        """
        print(f"Relative Humidity: {self.rh}")

JJs_old_sensor = Sensor_v1()
JJs_new_sensor = Sensor_v2()

print(help(JJs_old_sensor), end="\n\n")
print(help(JJs_new_sensor), end="\n\n")