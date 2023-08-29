from os import getenv

import psycopg2
from dotenv import load_dotenv
from psycopg2 import extras

load_dotenv()

db_config: dict[str, str] = {
    "database": getenv("DATABASE"),
    "host": getenv("HOST"),
    "user": getenv("SUPABASE_USERNAME"),
    "password": getenv("PASSWORD"),
    "port": getenv("PORT"),
}

connection = psycopg2.connect(**db_config)
db_cursor = connection.cursor(cursor_factory=psycopg2.extras.RealDictCursor)

if __name__ == "__main__":
    pass
