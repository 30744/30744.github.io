try:
    num = int(input("Enter a number: "))
    result = 10 / num
    print("Result:", result)
except:
    print("Division by zero is not allowed.")

try:
    num = int(input("Enter a number: "))
    result = 10 / num
    print("Result:", result)
except ZeroDivisionError:
    print("Error: Division by zero is not allowed.")

try:
    # Get inputs from the user
    company_name = input("Enter the insurance company name: ")
    num_employees = int(input("Enter the number of employees: "))
    location = input("Enter the location (city, country, or state): ")
    lowest_price = float(input("Enter the lowest price for a policy: "))
    highest_price = float(input("Enter the highest price for a policy: "))

    # Generate the output using an f-string
    output = (
        f"We are {company_name} located in {location}. Our {num_employees} employees can help you "
        f"find the policy that is right for you with policies ranging from ${lowest_price:.2f} to ${highest_price:.2f} per month!"
    )

    print(output)
except ValueError:
    print("Invalid input. Please enter a valid number for employees and policy prices.")
