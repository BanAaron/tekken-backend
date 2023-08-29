import json

from fastapi import FastAPI
from fastapi.responses import HTMLResponse

from db import db_cursor

app = FastAPI()


def sanitize(string: str) -> str:
    """
    strips and lowers strings
    :param string: any string
    :return: sanitized string
    """
    return string.strip().lower()


@app.get("/")
def read_root() -> HTMLResponse:
    html: str = """
     <html>
        <head>
            <title>Tekken 8 API</title>
        </head>
        <body>
            <h1>Tekken 8 API</h1>
            <p>Documentation can be found <a href="http://127.0.0.1:8000/docs">here</a></p>
        </body>
    </html>
    """
    return HTMLResponse(html)


@app.get("/character")
def get_characters() -> str:
    """
    gets data for all characters
    :return: json
    """
    query: str = f"""
        SELECT id, full_name, short_name, nationality, sex, age, weight_kg, height_cm
        FROM character
        ORDER BY id;
    """
    db_cursor.execute(query)
    result = db_cursor.fetchall()
    result_json = json.dumps(result)
    return result_json


@app.get("/character/{name}")
def get_character(name: str) -> str:
    """
    gets the provided characters data from the database
    :param name: characters first name as a string
    :return: id, full_name, short_name, nationality, sex, age, weight_kg, height_cm
    """
    name = sanitize(name)
    query = """
        SELECT id, full_name, short_name, nationality, sex, age, weight_kg, height_cm
        FROM character 
        WHERE short_name=%s
        """
    db_cursor.execute(query, (name,))  # (name,) creates a tuple with a single value
    result = db_cursor.fetchone()
    result_json = json.dumps(result)
    return result_json


if __name__ == "__main__":
    pass
