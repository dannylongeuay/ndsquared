from fastapi import FastAPI
from uuid import uuid4
from asyncio import sleep

app = FastAPI()


async def long_running_task(task_id: uuid4):
    for i in range(5):
        print(f'{task_id} {i}')
        await sleep(1)


@app.get('/')
async def index():
    task_id = uuid4()
    await long_running_task(task_id)
    return {'hello', 'world'}
