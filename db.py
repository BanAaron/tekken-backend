from os import getenv

import psycopg2
from dotenv import load_dotenv

load_dotenv()

DATABASE: str = getenv("DATABASE")
HOST: str = getenv("HOST")
USERNAME: str = getenv("SUPABASE_USERNAME")
PASSWORD: str = getenv("PASSWORD")
PORT: str = getenv("PORT")

connection = psycopg2.connect(
    database=DATABASE, host=HOST, user=USERNAME, password=PASSWORD, port=PORT
)

cursor = connection.cursor()
