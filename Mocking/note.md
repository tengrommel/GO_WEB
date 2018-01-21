# Mocking

1. Plan functionality and write test case
2. Run all tests and ensure that the new test fails.
    - If it passes, go back to Step 1 and adjust or extend tests.
    - If it fails, go to Step 3.
3. Implement planned functionality
4. Run all tests
    - If test(s) fails, go back to Step 3 and revise functionality implementation.
    - If all tests pass : 
        
        1.Refine code if necessary(and run tests)
        
        2.Continue with test case for next functionality(Step 1) unless all functionality implemented.
        

## Testing the Server/Invoked entity

- Approach: Writing tests against server functionality 
- Server = Model/Implementation
- Client = Test(Independent/dependent vars)

Challenge:
- Network slow
- Network hard to control

