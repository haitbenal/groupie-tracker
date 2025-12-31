# Groupie Trackers

## Objectives

**Groupie Trackers** is a web project that consists of receiving a given API and manipulating the data contained in it in order to create a user-friendly website that displays information about bands and artists.

The provided API is composed of four main parts:

1. **Artists**
   Contains information about bands and artists such as:

   * Name(s)
   * Image
   * Year they began their activity
   * Date of their first album
   * Members

2. **Locations**
   Contains information about the last and/or upcoming concert locations.

3. **Dates**
   Contains information about the last and/or upcoming concert dates.

4. **Relation**
   Links all the other parts together (artists, locations, and dates).

Using this data, the goal is to build a website that displays band and artist information through different data visualizations such as:

* Cards
* Lists
* Tables
* Pages
* Blocks
* Graphics (or any other visualization of your choice)

The choice of design and data presentation is left to the developer.

---

## Events & Client–Server Interaction

This project also focuses on the creation and visualization of **events/actions**.

An event (or action) is defined as a **client call to the server** (client–server interaction). This action must:

* Be triggered by the client (user interaction, time-based event, or any other factor)
* Communicate with the server
* Receive data using a **request–response** mechanism

More about request–response communication can be found here:
👉 [https://en.wikipedia.org/wiki/Request%E2%80%93response](https://en.wikipedia.org/wiki/Request%E2%80%93response)

---

## Instructions

* The backend **must be written in Go**.
* Only **standard Go packages** are allowed.
* The site and server **must not crash at any time**.
* All pages must work correctly.
* All possible errors must be handled properly.
* The code must follow **good coding practices**.
* It is **recommended** to include unit test files.

---

## Allowed Packages

* Only the **Go standard library** is allowed.

---

## Usage

### Running the project

```bash
git clone git@github.com:haitbenal/groupie-tracker.git
cd groupie-tracker
go run .
```

Then open your browser and visit:

```
http://localhost:8080
```

Make sure no other service is using the same port.

---

## Learning Outcomes

This project helps you learn and practice:

* Data manipulation and storage
* Working with JSON files and formats
* HTML templating
* Event creation and visualization
* Client–server architecture
* Building a web server using Go

---

## Troubleshooting

Something is wrong?

* Make sure the server is running correctly
* Check API availability and responses
* Verify routes and handlers
* Ensure proper error handling is implemented

---

## Author

Developed by:

* **Hichame Ait benalla**
* **Yaakoub Elhaouari**

As part of the **Groupie Trackers** project using Go.
