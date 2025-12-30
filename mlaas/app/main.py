from fastapi import FastAPI
from tensorflow.keras.models import load_model

from app.dto import PredictionRequest
from app.utils.converters import y_to_float, request_to_x

model = load_model("../model_artifacts/model.keras")
app = FastAPI()

@app.get("/")
async def root():
    return {"message": "Welcome to the MLaaS API"}

@app.post("/predict")
def predict(request: PredictionRequest):
    x = request_to_x(request)
    y = model.predict(x)

    return {"prediction": y_to_float(y)}
