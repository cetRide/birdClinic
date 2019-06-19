var x = []
var finalData
var languages = [

    {
        "language": "Bluish/ Gray eye",
        "value": 1
    },
    {
        "language": "Breathing problems",
        "value": 2
    },
    {
        "language": "Closed eyes",
        "value": 3
    },
    {
        "language": "Coughing",
        "value": 4
    },
    {
        "language": "Cyanosis",
        "value": 5
    },
    {

        "language": "Dehydration",
        "value": 6
    },
    {
        "language": "Depression",
        "value": 7
    },
    {

        "language": "Diarrhoea",
        "value": 8
    },
    {

        "language": "Discoloured eggs",
        "value": 9
    },
    {

        "language": "Droopy looking feathers",
        "value": 10
    },
    {

        "language": "Drop in egg production",
        "value": 11
    },
    {

        "language": "Drowsiness",
        "value": 12
    },
    {

        "language": "Dry unkempt feathers",
        "value": 13
    },
    {

        "language": "Dullness",
        "value": 14
    },
    {

        "language": "Enlarged feather follicles",
        "value": 15
    },
    {

        "language": "Facial swelling",
        "value": 16
    },
    {

        "language": "Greenish diarrhea",
        "value": 17
    },
    {

        "language": "Hanged head",
        "value": 18
    },
    {

        "language": "High mortality",
        "value": 19
    },
    {

        "language": "Hunched up",
        "value": 20
    },
    {

        "language": "Inactivity",
        "value": 21
    },
    {

        "language": "Inappetence",
        "value": 22
    },
    {

        "language": "Increased thirst",
        "value": 23
    },
    {

        "language": "Lack of appetite",
        "value": 24
    },
    {

        "language": "Lethargy",
        "value": 25
    },
    {

        "language": "Listless chicks",
        "value": 26
    },
    {

        "language": "Listlessness",
        "value": 27
    },
    {

        "language": "Mucoid diarrhea",
        "value": 28
    },
    {

        "language": "Nasal discharge",
        "value": 29
    },
    {

        "language": "Oral discharge",
        "value": 30
    },
    {

        "language": "Paralysis of legs",
        "value": 31
    },
    {

        "language": "Paralysis of neck",
        "value": 32
    },
    {

        "language": "Paralysis of wings",
        "value": 33
    },
    {

        "language": "Picking at own vent",
        "value": 34
    },

    {
        "language": "Picking up food then dropping it",
        "value": 35
    },
    {

        "language": "Poor egg albumen",
        "value": 36
    },
    {

        "language": "Poor growth",
        "value": 37
    },
    {

        "language": "Poor weight gain",
        "value": 38
    },
    {

        "language": "Rapid drop in feeding",
        "value": 39
    },
    {

        "language": "Rapid drop in water consumption",
        "value": 40
    },
    {

        "language": "Rapid weight loss",
        "value": 41
    },
    {

        "language": "Rattling",
        "value": 42
    },
    {

        "language": "Reduced appetite",
        "value": 43
    },
    {

        "language": "Reduced appetite",
        "value": 44
    },
    {

        "language": "Reduced appetite",
        "value": 45
    },
    {

        "language": "Reluctance to eat",
        "value": 46
    },
    {

        "language": "Ruffled feathers",
        "value": 47
    },
    {

        "language": "Ruffled feathers",
        "value": 48
    },
    {

        "language": "Ruffled feathers",
        "value": 49
    },
    {

        "language": "Runny nose",
        "value": 50
    },
    {

        "language": "Sharp drop in egg production",
        "value": 51
    },
    {
        "language": "Sitting in hinged position",
        "value": 52
    },
    {

        "language": "Sleeping with beak touching the floor",
        "value": 53
    },
    {

        "language": "Sneezing",
        "value": 54
    },
    {

        "language": "Sneezing",
        "value": 55
    },
    {

        "language": "Soft egg shell",
        "value": 56
    },
    {

        "language": "Soiled vent feathers",
        "value": 57
    },
    {

        "language": "Swollen wattles",
        "value": 58
    },
    {

        "language": "Unsteady gait",
        "value": 59
    },
    {

        "language": "Watery diarrhea",
        "value": 60
    },
    {

        "language": "Weakness",
        "value": 61
    },
    {

        "language": "Weight loss",
        "value": 62
    },
    {

        "language": "Weight loss",
        "value": 63
    },
    {
        "language": "White/greenish watery mucoid diarrhea",
        "value": 64
    }

];


var groupData = [
    {
        "groupName": "B",
        "groupData": [
            {
                "language": "Bluish/ Gray eye",
                "value": 1
            },
            {
                "language": "Breathing problems",
                "value": 2
            },
        ]
    },
    {
        "groupName": "C",
        "groupData": [
            {
                "language": "Closed eyes",
                "value": 3
            },
            {
                "language": "Coughing",
                "value": 4
            },
            {
                "language": "Cyanosis",
                "value": 5
            },
        ]
    },
    {
        "groupName": "D",
        "groupData": [
            {

                "language": "Dehydration",
                "value": 6
            },
            {
                "language": "Depression",
                "value": 7
            },
            {

                "language": "Diarrhoea",
                "value": 8
            },
            {

                "language": "Discoloured eggs",
                "value": 9
            },
            {

                "language": "Droopy looking feathers",
                "value": 10
            },
            {

                "language": "Drop in egg production",
                "value": 11
            },
            {

                "language": "Drowsiness",
                "value": 12
            },
            {

                "language": "Dry unkempt feathers",
                "value": 13
            },
            {

                "language": "Dullness",
                "value": 14
            },
        ]
    },
    {
        "groupName": "E",
        "groupData": [
            {

                "language": "Enlarged feather follicles",
                "value": 15
            },
        ]
    },
    {
        "groupName": "F",
        "groupData": [
            {

                "language": "Facial swelling",
                "value": 16
            },
        ]
    },
    {
        "groupName": "G",
        "groupData": [
            {

                "language": "Greenish diarrhea",
                "value": 17
            },
        ]
    },
    {
        "groupName": "H",
        "groupData": [
            {

                "language": "Hanged head",
                "value": 18
            },
            {

                "language": "High mortality",
                "value": 19
            },
            {

                "language": "Hunched up",
                "value": 20
            },
        ]
    },
    {
        "groupName": "I",
        "groupData": [
            {

                "language": "Inactivity",
                "value": 21
            },
            {

                "language": "Inappetence",
                "value": 22
            },
            {

                "language": "Increased thirst",
                "value": 23
            },
        ]
    },
    {
        "groupName": "L",
        "groupData": [
            {

                "language": "Lack of appetite",
                "value": 24
            },
            {

                "language": "Lethargy",
                "value": 25
            },
            {

                "language": "Listless chicks",
                "value": 26
            },
            {

                "language": "Listlessness",
                "value": 27
            },
        ]
    },
    {
        "groupName": "M",
        "groupData": [
            {

                "language": "Mucoid diarrhea",
                "value": 28
            },
        ]
    },
    {
        "groupName": "N",
        "groupData": [
            {

                "language": "Nasal discharge",
                "value": 29
            },

        ]
    },
    {
        "groupName": "O",
        "groupData": [
            {

                "language": "Oral discharge",
                "value": 30
            },
        ]
    },
    {
        "groupName": "P",
        "groupData": [
            {

                "language": "Paralysis of legs",
                "value": 31
            },
            {

                "language": "Paralysis of neck",
                "value": 32
            },
            {

                "language": "Paralysis of wings",
                "value": 33
            },
            {

                "language": "Picking at own vent",
                "value": 34
            },

            {
                "language": "Picking up food then dropping it",
                "value": 35
            },
            {

                "language": "Poor egg albumen",
                "value": 36
            },
            {

                "language": "Poor growth",
                "value": 37
            },
            {

                "language": "Poor weight gain",
                "value": 38
            },
        ]
    },
    {
        "groupName": "R",
        "groupData": [
            {

                "language": "Rapid drop in feeding",
                "value": 39
            },
            {

                "language": "Rapid drop in water consumption",
                "value": 40
            },
            {

                "language": "Rapid weight loss",
                "value": 41
            },
            {

                "language": "Rattling",
                "value": 42
            },
            {

                "language": "Reduced appetite",
                "value": 43
            },
            {

                "language": "Reduced appetite",
                "value": 44
            },
            {

                "language": "Reduced appetite",
                "value": 45
            },
            {

                "language": "Reluctance to eat",
                "value": 46
            },
            {

                "language": "Ruffled feathers",
                "value": 47
            },
            {

                "language": "Ruffled feathers",
                "value": 48
            },
            {

                "language": "Ruffled feathers",
                "value": 49
            },
            {

                "language": "Runny nose",
                "value": 50
            },
        ]
    },
    {
        "groupName": "S",
        "groupData": [
            {

                "language": "Sharp drop in egg production",
                "value": 51
            },
            {
                "language": "Sitting in hinged position",
                "value": 52
            },
            {

                "language": "Sleeping with beak touching the floor",
                "value": 53
            },
            {

                "language": "Sneezing",
                "value": 54
            },
            {

                "language": "Sneezing",
                "value": 55
            },
            {

                "language": "Soft egg shell",
                "value": 56
            },
            {

                "language": "Soiled vent feathers",
                "value": 57
            },
            {

                "language": "Swollen wattles",
                "value": 58
            },
        ]
    },
    {
        "groupName": "U",
        "groupData": [
            {

                "language": "Unsteady gait",
                "value": 59
            },
        ]
    },
    {
        "groupName": "W",
        "groupData": [
            {

                "language": "Watery diarrhea",
                "value": 60
            },
            {

                "language": "Weakness",
                "value": 61
            },
            {

                "language": "Weight loss",
                "value": 62
            },
            {

                "language": "Weight loss",
                "value": 63
            },
            {
                "language": "White/greenish watery mucoid diarrhea",
                "value": 64
            }

        ]
    },
];

var settings = {
    "inputId": "languageInput",
    "data": languages,
    "groupData": groupData,
    "itemName": "language",
    "groupItemName": "groupName",
    "groupListName": "groupData",
    "container": "transfer",
    "valueName": "value",
    "callable": function (data, names) {
        console.log("Selected ID：" + data)
        $("#selectedItemSpan").text(names)
        finalData = names
        var result = []
        var i
        for (i = 0; i < data.length; i++) {
            result[i] = data[i]
        }

    }
};
Transfer.transfer(settings);

$('#mySubmit').click(function () {
    console.log(finalData)

    axios.post('/symptom_submit', {
        data: finalData
    })
        .then(data => {
            window.location.href = data.data.Response
            alert("Symptoms successfully submitted")
        })
        .catch(err => {
            var respo = err.response.data

            if (respo.Results == "invalid") {
                alert("No disease is found") ///error page !!!
                window.location.href = data.data.Response

            }

        })
})