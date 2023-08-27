from fastapi import FastAPI
from db import db

app = FastAPI()


@app.get("/")
def read_root():
    return {"Hello", "World"}


@app.get("/character/{name}")
def get_character(character_name: str) -> dict[str, str]:
    result: str = db.execute(f"SELECT name FROM character WHERE name={character_name}")
    return {"name": result}
