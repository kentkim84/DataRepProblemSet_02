GMIT Software Development 3rd year

Module: Data Representation

GoLang Problem sheet

### Go Examples

This repository contains some example code written in the programming language Go. The author is Yongjin Kim, a Student at GMIT.

This problem set is to learn the fundamentals of creating a web application in Go. 

## How to clone
1. Open **Git Bash**.
2. Change the current working directory to the location where you want the cloned directory to be made.
3. Type 'git clone', and then paste the URL 'https://github.com/kentkim84/GoLangWebAppProject.git' or 'git@github.com:kentkim84/GoLangWebAppProject.git'
git clone https://github.com/kentkim84/GoLangWebAppProject.git

## Coding Standards
// Version 0.1 using C standards

### 1. Naming Conventions and Style
1.1. Use Pascal casing for class and structs
    
    class PlayerManager;
    struct AnimationInfo;

1.2. Use camel casing for local variable names and function parameters
    
    void SomeMethod(const int someParameter);
    {
        int someNumber;
    
    }

1.3. Use verb-object pairs for method names
a.	Use pascal casing for public methods
        
    public:
    void DoSomething();

b.	Use camel casing for other methods
        
    private:
    void doSomething();

### References
This examples are from
* [Data Representation and Querying - Ian McLoughlin](https://data-representation.github.io/problems/go-web-applications.html)

This coding standard is followed by
* [popeKim's c/c++ coding standards](https://docs.google.com/document/d/1cT8EPgMXe0eopeHvwuFmbHG4TJr5kUmcovkr5irQZmo/edit#heading=h.r2n9mhxbh2gg)