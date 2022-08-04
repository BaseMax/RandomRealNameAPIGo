# Random Real Name

## Assignment

### Technology

Java and Kotlin, Ktor (which is a server-side and client-side network library)

In most projects, we need to generate random names, but we do not want to use random characters. We need real names. 

We have an email provider which provides your email address. By getting this repository they will be able to generate many random email address but meanfull names.

**The challenge of this is to generate a meaningful but UNIQUE name!**

Suppose you are going to generate 10k email address, so you need to have a lot of unique names. An email address is a mix of name, family, age, or maybe a date.
So by using different value we are going to generate unique name.

### Communication way

We are going to use **JSON standard format** and so our API / web service will be RESTful.

### Routes

- `GET /get`

**Parameters:**
Default value of `limit` is 1 but you are able to change limit to any number of emails you need. like 100.

This route not going to generate an email for you. this only generates random and UNIQUE names.


