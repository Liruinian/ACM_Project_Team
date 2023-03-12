function parsejson(){
  json = document.getElementById("timetablejson").value
  parsed = JSON.parse(json)
  parsed = parsed.dateList[0].selectCourseList
  console.log(parsed)
  var newCourseList =[
    ['', '', '', '', '', '', '', '', '', '', '', ''],
    ['', '', '', '', '', '', '', '', '', '', '', ''],
    ['', '', '', '', '', '', '', '', '', '', '', ''],
    ['', '', '', '', '', '', '', '', '', '', '', ''],
    ['', '', '', '', '', '', '', '', '', '', '', ''],
    ['', '', '', '', '', '', '', '', '', '', '', ''],
  ];

  for (i in parsed) {
    classp = parsed[i]
    cteacher = classp.attendClassTeacher
    for (j in classp.timeAndPlaceList){
      coursearr = classp.timeAndPlaceList[j]
      cname = coursearr.coureName+"\n"+cteacher+"\n"+coursearr.teachingBuildingName+coursearr.classroomName+"\n"+coursearr.weekDescription
      cweek = coursearr.classDay
      ctime = coursearr.classSessions
      cdur = coursearr.continuingSession
      for(let k=0; k<cdur; k++){
        newCourseList[cweek-1][ctime+k-1] = cname
      }
    }
  }
  console.log(newCourseList)

  Timetable.setOption({
    timetables: newCourseList,
    timetableType: courseType,
  });
}


var courseList = [
  ['', '', '', '', '', '', '', '', '', '', '', ''],
  ['', '', '', '', '', '', '', '', '', '', '', ''],
  ['', '', '', '', '', '', '', '', '', '', '', ''],
  ['', '', '', '', '', '', '', '', '', '', '', ''],
  ['', '', '', '', '', '', '', '', '', '', '', ''],
  ['', '', '', '', '', '', '', '', '', '', '', ''],
];
  
  var day = new Date().getDay();
  var courseType = [
    [{index: '1', name: '08:10\n08:55'}, 1],
    [{index: '2', name: '09:00\n09:45'}, 1],
    [{index: '3', name: '10:05\n10:50'}, 1],
    [{index: '4', name: '10:55\n11:40'}, 1],
    [{index: '5', name: '13:30\n14:15'}, 1],
    [{index: '6', name: '14:20\n15:05'}, 1],
    [{index: '7', name: '15:35\n16:20'}, 1],
    [{index: '8', name: '16:25\n17:10'}, 1],
    [{index: '9', name: '18:30\n19:15'}, 1],
    [{index: '10', name: '19:20\n20:05'}, 1],
    [{index: '11', name: '20:15\n21:00'}, 1],
    [{index: '12', name: '21:05\n21:50'}, 1]
  ];


