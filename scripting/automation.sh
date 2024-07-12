#!/bin/bash

# Run data processing script
Rscript ../data_processing/process_data.R

# Train machine learning model
python3 ../machine_learning/model_training.py

# Start backend server
cd ../backend
go run main.go &
