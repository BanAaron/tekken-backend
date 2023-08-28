from fastapi import FastAPI

from db import cursor

app = FastAPI()


def sanitize(string: str) -> str:
    """
    strips and lowers strings
    :param string: any string
    :return: sanitized string
    """
    string = string.strip().lower()
    return string


@app.get("/")
def read_root():
    return {"Hello", "World"}


@app.get("/character")
def get_characters():
    """gets data for all characters
    :return: id, full_name, short_name, nationality, sex, age, weight_kg, height_cm
    """
    query: str = f"""
        SELECT id, full_name, short_name, nationality, sex, age, weight_kg, height_cm
        FROM character
        order by id;
    """
    cursor.execute(query)
    result = cursor.fetchall()
    return result


@app.get("/character/{name}")
def get_character(name: str):
    """gets the provided characters data from the database
    :param name: characters first name as a string
    :return: id, full_name, short_name, nationality, sex, age, weight_kg, height_cm
    """
    name = sanitize(name)
    query = """
        SELECT id, full_name, short_name, nationality, sex, age, weight_kg, height_cm
        FROM character 
        WHERE short_name=%s
        """
    cursor.execute(query, (name,))  # (name,) creates a tuple with a single value
    result = cursor.fetchone()
    return result


if __name__ == "__main__":
    get_character()
