from dataclasses import dataclass

import nats
from nats.aio.client import Client

import logging
import json


@dataclass
class NatsClient:
    __client: Client | None = None

    async def connect(self, address: str):
        self.__client = await nats.connect(address)
        logging.getLogger().info(f"✅ Connected to NATS broker at {address}")

    async def subscribe(self, subject: str):
        await self.__client.subscribe(subject, cb=NatsClient.__message_handler)
        logging.getLogger().info(f"NATS Client subscribed to subject: {subject}")

    async def disconnect(self):
        if self.__client:
            logging.getLogger().info("Disconnecting NATS client")
            await self.__client.drain()
            await self.__client.close()

    @staticmethod
    async def __message_handler(msg):
        subject = msg.subject
        reply = msg.reply

        payload = json.loads(msg.data.decode())
        pretty = json.dumps(payload, indent=2)
        print("Received a message on '{subject} {reply}': {data}".format(
            subject=subject, reply=reply, data=pretty)
        )
