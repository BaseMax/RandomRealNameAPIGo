# Random Real Name

A short description of the project. TODO....

## Assignment

In most projects, we need to generate random names, but we do not want to use random characters. We need real names. 

We have an email provider which provides your email address. By getting this repository they will be able to generate many random email address but meanfull names.

**The challenge of this is to generate a meaningful but UNIQUE name!**

Suppose you are going to generate 10k email address, so you need to have a lot of unique names. An email address is a mix of name, family, age, or maybe a date.
So by using different value we are going to generate unique name.

### Communication way

We are going to use **JSON standard format** and so our API / web service will be RESTful.

### Technology

Java and Kotlin, Ktor (which is a server-side and client-side network library)

### Routes

- `GET /`

A route that will show this document and everyone can study more about the project as well as this README.markdown.

- `GET /get`

**Parameters:**
Default value of `limit` is 1 but you are able to change limit to any number of emails you need. like 100.

This route not going to generate an email for you. this only generates random and UNIQUE names.

**Output:**

```json
{
   "status": 1,
   "name": "alireza2000"
}
```

Or maybe as following if you are asking for more than one name:

```json
{
   "status": 1,
   "names": [
    "alireza2004",
    "hamid.h3000"
    "max.base1"
   ]
}
```

Or if something went wrong:

```json
{
   "status": 0,
   "message": "Oops, sorry. Something does not go as we expected."
}
```

**Note:** Obviously, the HTTP Code is expected to be 200 if the request is answered successfully. Otherwise, you can use the appropriate code according to the error.

## Database or not?

As it is clear, maybe there is no reason to use database in this project. But you can use the database you like.

You can also use the file to store names and surnames.
But keep in mind that your web service must be able to respond to a large number of simultaneous requests.

## Authors

- Amir Shiati
- Max Base

Anyone is welcome to contribute, change or develop this project. Thanks in advance, Any comments are appreciated.
