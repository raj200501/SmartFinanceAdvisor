library(dplyr)

# Load dataset
data <- read.csv('raw_financial_data.csv')

# Process data
processed_data <- data %>%
  mutate(balance = income - expenses) %>%
  select(user_id, balance, income, expenses)

# Save processed data
write.csv(processed_data, 'financial_data.csv', row.names = FALSE)

print("Data processed and saved.")
