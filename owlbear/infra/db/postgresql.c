#include <libpq-fe.h>
#include <stdio.h>
#include <stdlib.h>

PGconn *conn() {
  // Connect to the database
  // conninfo is a string of keywords and values separated by spaces.
  char *conninfo = "dbname=your_db_name user=your_user_name "
                   "password=your_password host=localhost port=5432";

  // Create a connection
  PGconn *conn = PQconnectdb(conninfo);

  // Check if the connection is successful
  if (PQstatus(conn) != CONNECTION_OK) {
    // If not successful, print the error message and finish the connection
    printf("Error while connecting to the database server: %s\n",
           PQerrorMessage(conn));

    // Finish the connection
    PQfinish(conn);

    // Exit the program
    exit(1);
  }

  // We have successfully established a connection to the database server
  printf("Connection Established\n");
  printf("Port: %s\n", PQport(conn));
  printf("Host: %s\n", PQhost(conn));
  printf("DBName: %s\n", PQdb(conn));

  // Close the connection and free the memory
  // PQfinish(conn);

  return conn;
}
