# contactAPI

## Routes 

### Endpoints

[GET] <code>"/api/contacts" </code> Get All Contacts <br>
<br>
[GET] <code>"/api/contacts/{id}"</code>  Get a single contact, pass id as argument in URL <br>
<br>
[POST] <code>"/api/contacts" </code> Adds a new contact, post contact(name, contact, email) as json <br>
Example
<code>
{
    "name": "Aryan Maurya",
    "contact": "9971984993, 9795644147",
    "email": "aryanmaurya1@outlook.com"
}
</code>
<br>
<br>
[PUT] <code>"/api/contacts/{id}" </code> Updates an existing contact based on the id passed in URL as argument, pass complete updated contact even if updating a single field as json.
<br>
<br>
[DELETE] <code>"/api/contacts/{id}" </code> Deletes an 	existing contact based on the id passed in URL as argument
