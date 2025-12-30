from pathlib import Path

from fastapi import FastAPI
from tensorflow.keras.models import load_model

from app.dto import PredictionRequest
from app.utils.converters import y_to_float, request_to_x

BASE_DIR = Path(__file__).resolve().parent.parent
MODEL_PATH = BASE_DIR / "model_artifacts" / "model.keras"

model = load_model(MODEL_PATH)
app = FastAPI()

@app.get("/")
async def root():
    return {"message": "Welcome to the MLaaS API"}

@app.post("/predict")
def predict(request: PredictionRequest):
    x = request_to_x(request)
    y = model.predict(x)

    return {"prediction": y_to_float(y)}
