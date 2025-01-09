# Pokedex CLI  

Welcome to the **Pokedex CLI**, a gamified command-line interface built with **Go (Golang)**! 🎮🐾  

This tool allows you to interact with the Pokédex API in a fun and engaging way. You can explore locations, catch Pokémon, inspect their details, and maintain your personal Pokédex—all from your terminal!  

## Features  

- **Catch Pokémon**: Find and catch Pokémon by name.  
- **Explore Locations**: Navigate through different locations to discover Pokémon.  
- **Inspect Pokémon**: View detailed stats and information about the Pokémon you've caught.  
- **Maintain a Pokedex**: Keep track of all the Pokémon you've captured.  
- **Caching**: Improved performance with cached data to reduce redundant API calls.  

## Commands  

- **`catch <pokemon_name>`**: Catch a Pokémon by its name.  
- **`map`**: Get the next page of locations.  
- **`mapb`**: Get the previous page of locations.  
- **`explore <location_name>`**: Explore a specific location to discover Pokémon.  
- **`inspect <pokemon_name>`**: Inspect a Pokémon to view its stats and details.  
- **`pokedex`**: Display all the Pokémon you’ve caught.  
- **`help`**: Display a help message with available commands.  
- **`exit`**: Exit the Pokedex CLI.  

## Installation  

1. Clone this repository:  
   ```bash  
   git clone https://github.com/hawkaii/pokedexcli.git  
   ```  

2. Navigate to the project directory:  
   ```bash  
   cd pokedexcli  
   ```  

3. Build the CLI tool:  
   ```bash  
   go build -o pokedex  
   ```  

4. Run the CLI:  
   ```bash  
   ./pokedex  
   ```  

## Requirements  

- Go 1.19 or later.  
- Internet connection for API requests.  

## Future Enhancements  

- Save and load progress between sessions.  
- Add difficulty levels for catching Pokémon.  
- Introduce a battle system between caught Pokémon.  
- Improve CLI aesthetics and interactivity.  

## Contributing  

Contributions are welcome! If you’d like to contribute, feel free to fork the repository, make your changes, and submit a pull request.  

