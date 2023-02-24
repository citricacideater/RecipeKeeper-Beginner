# UNDER CONSTRUCTION

# RecipeKeeper-Beginner
Code for making a static online recipe keeper using Go

## Final Base Code
This folder contains the finished end product.
It is a REST api? that takes a given JSON as an input (https://api.npoint.io/fff0f131782057b16a12) and returns the recipes as 'cards' at the root ("/"). Selecting a recipe card takes you to another page ("/recipe/[recipeId]") that shows the full information of the recipe.
Data is fed into the pages through the use of the html/template package in the Go standard library.
For this tutorial the website layout is designed in CSS using CSS Grid, however once completed feel free to play with design and layout.

## Code Steps
This folder contains subfolders breaking the code into steps. Each step contains an instructions.txt file detailing what is happening/what you should do.
When following the steps don't create seperate folders/files per step.

## Bonus Features
This folder contains variations on the 'Base Code' with additional features, such as searching by tags.
