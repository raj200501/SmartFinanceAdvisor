import unittest
from sklearn.linear_model import LinearRegression
import joblib

class TestModelTraining(unittest.TestCase):
    def test_model_training(self):
        model = joblib.load('financial_model.pkl')
        self.assertIsInstance(model, LinearRegression)

if __name__ == '__main__':
    unittest.main()
