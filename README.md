# AirBnb For Stuff 📦

Our application is a social platform that connects item owners (lenders) with renters, enabling peer-to-peer rental of personal belongings—like Airbnb, but for everyday items. 
Users can log in to browse a wide range of available items, rent what they need, and complete payments directly through the app. Owners can list items they'd like to lend and also rent from others, creating a flexible, two-way marketplace. 

Docs: [Google Docs](https://docs.google.com/document/d/1pyazoKCPFO2WeYyncGOaY4HedAxXi0jh/edit)

Built with [Svelte](https://svelte.dev/), [GO](https://go.dev/), [PostgresSQL](https://www.postgresql.org/).

## App Structure 🧱

```
├── backend
│   ├── main.go
|   ├── go.mod
|   ├── go.sum
|   ├── .env
│   ├── db                  // for DB connections
│   │   ├── db.go.go
│   │   ├── queries.sql
│   │   ├── schema.sql      // DB schema
│   ├── handlers          // API core handlers
│   │   ├── handlers.go
|   ├── middleware          // auth, CORS
│   │   ├── auth.go
│   │   ├── cors.go
│   └── models
│   |    └── model_functions.go     // DB functions
│   |    └── ...                    // Models for our application
|   |
├── frontend    //svelte app
├── data    // mock data generation
└── backup  // application's data
```

## Run Svelte ⚡️

```
cd frontend
npm i # install dependencies
npm run dev 
npm run build # production build
npm run preview # preview production build
```

Check `localhost:5173/`

## Run GO 💻
 
Create a `.env` file in `/backend`:

```
POSTGRES_URL=postgres://<username>:<password>@<host>:<port>/<dbname>?sslmode=require
```

Run the program

```
cd backend
go mod tidy # install dependencies
go run main.go 
```

Check `localhost:8080/`

## Generate mock data 📊

```
cd data
python -m venv venv 
source venv/bin/activate # activate virtual env
pip install -r requirements.txt
python script.py
```

Mock data should be generated in `/fake_data_csv`

###  Backup data are stored in `data/` 💿
