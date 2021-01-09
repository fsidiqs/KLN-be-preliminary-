Example of Code Smells wirtten in PHP Language  commonly found is DRY (Don't Repeat Yourself).
in bad.php file we can see the method `set_settings` from AwesomeAddon class share the same functionality with the method `set_settings` from EvenMoreAwesomeAddon class.
The solution for this is to create an abstract class which holds the common functionality and make the other two classess extend the abstract class. 
The solution can be seen in good.php file