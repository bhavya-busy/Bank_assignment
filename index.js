const express=require('express');
const app=express();  

app.use(express.json());
const bodyParser = require("body-parser")

const cors = require("cors")
const BankRoutes = require("./NodeRoutes/bankroute")
//const BranchRoutes = require("./node-routes/branch.route")


app.use(cors())
app.use(bodyParser.json())

app.use("/bank" , BankRoutes)

app.listen(8080, () => {
    console.log("listening on port 8080")
})