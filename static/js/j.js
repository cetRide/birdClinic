var x = []
var finalData
var languages = [
    {
        "language": "jQuery",
        "value": 122
    },
    {
        "language": "AngularJS",
        "value": 132
    },
    {
        "language": "ReactJS",
        "value": 422
    },
    {
        "language": "VueJS",
        "value": 232
    },
    {
        "language": "JavaScript",
        "value": 765
    },
    {
        "language": "Java",
        "value": 876
    },
    {
        "language": "Python",
        "value": 453
    },
    {
        "language": "TypeScript",
        "value": 125
    },
    {
        "language": "PHP",
        "value": 633
    },
    {
        "language": "Ruby on Rails",
        "value": 832
    }
];

var groupData = [
    {
        "groupName": "JavaScript",
        "groupData": [
            {
                "language": "jQuery",
                "value": 122
            },
            {
                "language": "AngularJS",
                "value": 643
            },
            {
                "language": "ReactJS",
                "value": 422
            },
            {
                "language": "VueJS",
                "value": 622
            }
        ]
    },
    {
        "groupName": "Popular",
        "groupData": [
            {
                "language": "JavaScript",
                "value": 132
            },
            {
                "language": "Java",
                "value": 112
            },
            {
                "language": "Python",
                "value": 124
            },
            {
                "language": "TypeScript",
                "value": 121
            },
            {
                "language": "PHP",
                "value": 432
            },
            {
                "language": "Ruby on Rails",
                "value": 421
            }
        ]
    }
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
        .then(function (response) {
            console.log(response);
        })
        .catch(function (error) {
            console.log(error);
        });

});