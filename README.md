### Run the server

```Go
cd backend
go mod tidy # install depencies
go run main.go
```

Check `localhost:8080/healthcheck`

### Generate mock data

```python
cd mock-data
python -m venv venv
source venv/bin/activate
pip install -r requirements.txt
python script.py
```

Mock data should be generated in `/fake_data_csv`
