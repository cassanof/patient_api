import json
from flask import Flask, request
from flask_restful import Resource, Api
from flask_cors import CORS
from tensorflow import keras
import numpy as np
from sklearn.preprocessing import StandardScaler
from joblib import dump, load

app = Flask(__name__)
CORS(app)
api = Api(app)

scaler = load('std_scaler.bin')
model = keras.models.load_model('probModel')


class Prediction(Resource):

    def post(self):
        json_data = request.get_json(force=True)
        raceStr = json_data['ethnicity']
        index = json_data['index_rel']
        ethIndex = 0
        try:
            if "white" in raceStr.lower():
                ethIndex = 0
            elif "cau" in raceStr.lower():
                ethIndex = 0
            elif "black" in raceStr.lower():
                ethIndex = 1
            elif "afri" in raceStr.lower():
                ethIndex = 1
            elif "hispan" in raceStr.lower():
                ethIndex = 2
            else:
                ethIndex = 3
        except:
            ethIndex = 0
        ethIndex = float(ethIndex)
        pred = np.argmax(model.predict(
            np.array(scaler.transform([[ethIndex, index]])))[0])
        pred = pred.item()
        return {'prediction': pred}


api.add_resource(Prediction, '/api')


if __name__ == '__main__':
    app.run(threaded=True, port=5000)
