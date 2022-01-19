# ASSIGNMENT- 1

## Run Local

To run the application locally,
- `cd item-inventory`
- `make run` 

### Topics Covered: Go I/O, String.  
   
Write a go program that accepts details (item name, item type, item prize) of different items from
Command line and outputs the item name, item prize, sales tax liability per item, final prize (sales tax + item prize) to the console. The input should be accepted with following command line options:

-name \<first item name>
-price \<price of first item>
-quantity \<quantity of first item>
-type \<type of first item>
 
The following functionalities/behavior is expected:

- All options other than item name can come in any order i.e. after -name you can have -price, -type option. Item type is a mandatory option. 
- The system should provide functionality to add more than one items i.e. Once the details of the first item is entered it should display a message saying:
    - Do you want to enter details of any other item (y/n):
    -  Appropriate behavior is required to be implemented as per user input for this question.
 
- Make use of go’s object oriented capabilities for implementing this business logic.
- Item type can have 3 possible values: raw, manufactured and imported.
- Tax rules for the 3 types are as follows:
    - raw: 12.5% of the item cost
    - manufactured: 12.5% of the item cost + 2% of (item cost + 12.5% of the item cost)
    - imported: 10% import duty on item cost + a surcharge (surcharge is: Rs. 5 if the final cost after applying tax & import duty is up to Rs. 100, Rs. 10 if the cost exceeds 100 and up to 200 and 5% of the final cost if it exceeds 200).
	
### Key Points:
 
- Use Go’s I/O capabilities to accept input from users.
- Use Go’s String functionalities to parse input strings.
- Coding conventions should be followed.
- Proper validation / info messages should be thrown on console.
- Do appropriate error handling wherever required.
- Where ever required please write comments in the code to make it more understandable.
- TDD methodology should be used

# ASSIGNMENT-2   

## Run Local

To run the application locally,
- `cd user-management`
- `make run` 
 
### Topics Covered: Go Serialization, Sorting.  
   
Write a menu driven command line go program that provides the following menu options:
- Add User details.
- Display User details.
- Delete User details
- Save User details.
- Exit
 
The option (1) should ask for the following user details. All the following details are mandatory and the program should perform the required validations (like blank data, integer value for age, roll number etc). Roll Number is a key to identify the uniqueness among the students.
- Full Name
- Age
- Address
- Roll Number
- Set of courses he is interested to enroll. There are a total of 6 courses (Course A, B, C, D, E and F). It is mandatory for each student to choose 4 out of 6 courses.
 
Once the validations are passed the user details should be added to an in memory data structure. The data structure should always keep the records sorted in ascending order. By default the records should be sorted on full name. If the name is the same for two students then sorting should be based on the roll number.
 
The option (2) should display the user details in the following format. Also the user should be provided with an option to sort the results (either in ascending or descending order) based on name, roll number, age, address.

- Name | RollNo | Age | Address | Courses  
 

The option (3) should ask for roll number and delete the student details corresponding to that roll number. Throw a proper user friendly message in case the roll number entered by the user does not exist.
 
The option (4) should save the in memory details of all the users to a disk. Use go’s serialization capabilities to serialize the in memory data to disk. If the user terminates the program after choosing this option the user’s data should be saved to disk and next time the user runs the program the in-memory collection should be pre populated with data already stored on the disk. 
 
The option (5) should terminate the program but before termination it should ask the user if he wants to save his latest changes (additions, deletions of users) to disk.
 
### Key Points:

- Use Go’s serialization mechanism to save user details to disk.
- Use Go proper pkgs/libs  for sorting.
- Coding conventions should be followed.
- Proper validation / info messages should be thrown on console.
- Student Info, course info, serialization code and command line menu code should be encapsulated in separate independent java classes.
- Where ever required please write comments in the code to make it more understandable.
- TDD methodology should be used
