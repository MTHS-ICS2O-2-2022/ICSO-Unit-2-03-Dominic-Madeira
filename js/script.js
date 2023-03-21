// Copyright (c) 2023 Dominic M. All rights reserved
//
// Created by: Dominic M.
// Created on: Sep 2020
// This file contains the JS functions for index.html

function myButtonClicked() {
  const houseNumber = document.getElementById("house-number").value
  const streetName = document.getElementById("street-name").value
  document.getElementById("address").innerHTML = "Your address is: " + houseNumber + " " + streetName
}
