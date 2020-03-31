# forms

Contains form validation code

 - Any code that has validation scopes which are poorly defined, especially dates, must be tested with sanity checks.
   - Dates for symptoms cannot have an onset in the future
   - Birthdates cannot be older than 117 years
   - binary choices, such as true false, must test to show other values cannot be entered
