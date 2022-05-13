db.createCollection('users', {
    validator: {
      $jsonSchema: {
        bsonType: 'object',
        required: ['first_name', 'last_name', 'birth_date', 'phone_number', 'email', 'roles'],
        properties: {
          first_name: {
            bsonType: 'string',
            description: 'first name of type string is required'
          },
          last_name: {
            bsonType: 'string',
            description: 'last name of type string is required'
          },
          birth_date: {
            bsonType: 'date',
            description: 'birth date of type date is required'
          },
          phone_number: {
            bsonType: 'string',
            description: 'phhone number of type string with length 10-15 is required',
            minLength: 10,
            maxLength: 15
          },
          email: {
            bsonType: 'string',
            description: 'email of type string is required'
          },
          roles: {
            bsonType: ['array'],
            minItems: 1,
            description: 'must be an array and is required',
            items:  {
                enum: [
                    {id: 1, name:'ADMIN'},
                    {id: 2, name:'CUSTOMER'},
                    {id: 3, name:'OWNER'}
                ]
            }
          }
        }
      }
    }
})


db.createCollection('business', {
  validator: {
    $jsonSchema: {
      bsonType: 'object',
      required: ['owner_id', 'name', 'address', 'phone_number', 'description', 'product_ids'],
      properties: {
        owner_id: {
          bsonType: 'objectId',
          description: 'owner id of type objectId is required'
        },
        name: {
          bsonType: 'string',
          description: 'name of type string is required'
        },
        address: {
          bsonType: 'string',
          description: 'address of type string is required'
        },
        phone_number: {
          bsonType: 'string',
          description: 'phhone number of type string with length 10-15 is required',
          minLength: 10,
          maxLength: 15
        },
        description: {
          bsonType: 'string',
          description: 'description of type string is required'
        },
        product_ids: {
          bsonType: ['array'],
          minItems: 1,
          description: 'must be an array and is required',
          items:  {
            bsonType: 'objectId'
          }
        }
      }
    }
  }
})

db.business.insertOne({
  owner_id: ObjectId("627cfe96b43286ade6b6e7f7"),
  name: "Faz's Cookies",
  address: 'Lengkong, Garawangi, Kuningan',
  phone_number: '082295106366',
  description: 'This is a cookie shop',
  photo_URL: 'http://shop-photo.com',
  products: [
    {
      name: 'cookie 1',
      price_IDR: Long("1000"),
      description: 'cookie number 1',
      photo_URL: 'http://cookie-1.com'
    },
    {
      name: 'cookie 2',
      price_IDR: Long("2000"),
      description: 'cookie number 2',
      photo_URL: 'http://cookie-2.com'
    },
    {
      name: 'cookie 3',
      price_IDR: Long("3000"),
      description: 'cookie number 3',
      photo_URL: 'http://cookie-3.com'
    }
  ]
})

db.createCollection('products', {
  validator: {
    $jsonSchema: {
      bsonType: 'object',
      required: ['name', 'price_IDR', 'description'],
      properties: {
        name: {
          bsonType: 'string',
          description: 'name of type string is required'
        },
        price_IDR: {
          bsonType: 'long',
          description: 'address of type string is required'
        },
        description: {
          bsonType: 'string',
          description: 'description of type string is required'
        }
      }
    }
  }
})

db.business.aggregate([
  {$match: {_id: ObjectId("627e7792e5dfc249b676c802")}},
  {$lookup:{from: 'products', localField:'product_ids', foreignField:'_id', as:'products'}}
])