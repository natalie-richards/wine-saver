# wine-saver# Wine Saver

Welcome to the Wine Saver project! This application helps you keep track of wines you like/dislike, ensuring you optimize your grocery trips and Costco runs.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Scope](#scope)

## Introduction

Wine Saver is designed as a central place to store memorable wines with photos and tasting notes.

## Features

- **Wine Photo Storage**: Save all your wine label photos
- **Label Identification**: An uploaded photo willl interact with a 3rd patrty API to return a list of label results
- **Wine Bookmarks**: Save this data in an easily accessible bookmark library

## Installation

To install Wine Saver, follow these steps:

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/wine-saver.git
   ```
2. Navigate to the project directory:
   ```sh
   cd wine-saver
   ```

### UI Setup

1. Navigate to the ui directory:

   ```sh
   cd ui
   ```

2. Install the UI dependencies:

   ```sh
   npm install
   ```

3. Create a .env file in the current directory:

   ```sh
   touch .env
   ```

4. Ask the author for the Rapid API credentials to paste here

5. Navigate to the gcs directory

   ```sh
   cd src/lib/gcs
   ```

6. Create a file 'service-account-key.json'. Ask the author for the service account key

### Backend Setup

1. Navigate to the backend directory from the root:

   ```sh
   cd backend
   ```

2. Create a .env file in the current directory. Ask the author for the database credentials
   ```sh
   touch .env
   ```

## Usage

To start using Wine Saver UI, run the following command from `ui`:

```sh
npm start
```

To start the server, run the following command from `backend`:

```sh
./run.sh
```

Open your browser and navigate to `http://localhost:5173` to access the application.

## Scope

The following is currently out of scope for this project:

- Auth - The app is currently designed for a single user
- Hosting - ask for access keys
- Front-end automated testing
