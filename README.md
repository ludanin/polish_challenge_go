# Goals

  - Create from 5 to 15 random "people" to walk on the map. To make our
  lives easier they're all going to be centered toward a specific coordinate,
  i.e. `lat = 59.955413` and `lng = 30.337844`.

  - A `Person` is a struct (or JavaScript object) that has a `ID`, a `Name`
  and a `Coordinate`

  - Those "people" must have their position updated at every second.

  - We need a server with a single route, it returns all the "people" available
  on the map.

# Completed

  - We would need to generate some random decimal noise to move these "people"
  around this location. For example: `lat += 0.00random` and `lng += 0.00random`

  - A `Coordinate` holds latitude and longitude values.
