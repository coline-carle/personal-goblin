import psycopg2
import progressbar
import logging
import argparse
import requests
import os

api_entrypoint = "eu.api.battle.net"

def main():
    args = parse_args()
    try:
        conn = psycopg2.connect(dbname="postgres", user="postgres", password="postgres", host="127.0.0.1")
    except:
        logging.error("can't connect to the database")
        os.exit(-1)
    insert_items(conn, "data/items.sql")
    update_items(conn, args.apikey)

def parse_args():
    parser = argparse.ArgumentParser(description='dump all world of warcraft items in a sqlite database')
    parser.add_argument('--database', action='store', dest='database', default='dbname=postgres user=postgres', help='sqlite database connect string')
    parser.add_argument('--apikey', action='store', dest='apikey', required=True, help='battle.net API key')
    return parser.parse_args()

def insert_items(conn, filename):
    with open(filename) as f:
        sql = f.read()
        cur = conn.cursor()
        cur.execute(sql)
        conn.commit()





def fetch_item(session, apikey, id):
  url = "https://{}/wow/item/{}?apikey={}".format(api_entrypoint, id, apikey)
  r = session.get(url)
  if r.status_code != 200:
    logging.error("fetching item", id, "->", r.status_code)
  return r.status_code, r.text

def update_items(conn, apikey):
    session = requests.Session()
    curg = conn.cursor()
    curg.execute("SELECT id FROM items")
    rows = curg.fetchall()
    curg.close()

    with progressbar.ProgressBar(max_value=len(rows)) as bar:
      for i, row in enumerate(rows):
        id = row[0]
        status, data = fetch_item(session, apikey, id)

        cur = conn.cursor()

        if status == 200:
          cur.execute("UPDATE items SET bnet_data = '{}' WHERE id = {}".format(data.replace("'", "''"), id))

        conn.commit()
        bar.update(i)

    conn.close()



if __name__ == "__main__":
  main()
