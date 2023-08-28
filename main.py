from fastapi import FastAPI

from db import cursor

app = FastAPI()


@app.get("/")
def read_root():
    return {"Hello", "World"}


@app.get("/character")
def get_characters():
    cursor.execute(f"SELECT character_name FROM characters")
    result = cursor.fetchall()
    return result


@app.get("/character/{name}")
def get_character(name: str):
    """ Gets character data from the database
    :param name: Character name as a string
    :return: Character data
    """
    cursor.execute(f"SELECT character_name FROM characters WHERE character_name='{name}'")
    result = cursor.fetchone()[0]
    return result


if __name__ == '__main__':
    get_character()
