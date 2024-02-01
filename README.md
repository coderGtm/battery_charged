# Battery_Charged

## Description
Battery_Charged is a simple tool to monitor HIGH Charge situations for Laptops. Generally Laptops don't have a mechanism to notify users when battery is charged. Continuing to charge post this can damage battery life, hence I created this utility to notigy me once the Battery reaches a certain level.

## Installation
Go should be installed on your system.

Simply Download the code and place the folder somewhere in your files where you store scripts. Tweak values if necessary (like the time intervals or battery file). You can then build the tool using:
```
go build
```
The `battery_charged` is an executable file that you now use. Placing in it boot script can automatically run it everytime you boot your laptop.

## Usage
You don't have to do anything after starting the executable. It keeps running in the background, consuming minimal amount of resources. (on my laptop, it was 0.0% CPU and 0.2% Memory)

## Contributing
You are welcome to contribute to this utility. Please open an issue first to discuss the modification before opening a PR.

## License
This project is licensed under the [MIT License](/LICENSE).