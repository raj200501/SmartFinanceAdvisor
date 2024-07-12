library(testthat)

test_that("Data processing works correctly", {
    data <- read.csv('financial_data.csv')
    expect_true(all(c('user_id', 'balance', 'income', 'expenses') %in% colnames(data)))
})

test_check("Data processing works correctly")
