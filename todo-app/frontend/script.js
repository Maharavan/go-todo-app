async function loadTask() {
    const getResponse = await fetch('http://localhost:3001/gettask')
    if (!getResponse.ok){
        throw new Error(`Error loading task`)
    }
    const data = await getResponse.json();
    renderTasks(data.tasks)
}
window.addEventListener('DOMContentLoaded',loadTask);
async function providenewTask(){
    try{
        const val = document.getElementById("taskdata").value;
        const postresponse = await fetch('http://localhost:3001/postnewtask', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({task: val,status:false})
            });
        
        if (!postresponse.ok){
            throw new Error(`POST request failed with status ${postresponse.status}`)
        }

        const postData = await postresponse.json()
        console.log('Post data successful')
        
        
        const getResponse = await fetch ('http://localhost:3001/gettask');

        if (!getResponse.ok){
            throw new Error(`GET request failed with status: ${getResponse.status}`);
        }

        const getData = await getResponse.json();
        console.log('GET request successful:', getData);
        await loadTask();
    }
    catch(error){
        console.error('Error found ',error)
    }

}

function renderTasks(tasks) {
    const taskList = document.getElementById("tasklist");
    taskList.innerHTML = ""; 

    tasks.forEach(taskout => {

        const main = document.createElement('div');
        main.classList.add('task-align'); 

        const label = document.createElement("label");
        label.htmlFor = taskout.id;
        label.textContent = taskout.task;

        const out = document.createElement("input");
        out.type = 'checkbox';
        out.id = taskout.id;
        out.name = "task-align"
        out.value = taskout.task
        out.checked = taskout.status;
        
        out.addEventListener("change",() =>{
            if (out.checked==true){
                main.style.backgroundColor = '#ffe7e9'
            }
            else{
                main.style.backgroundColor = '#ffffff'
            }
            updateTask(taskout.id,taskout.task,out.checked);
        });
        main.appendChild(label);
        main.appendChild(out);

        taskList.appendChild(main);
    });
    taskList.scrollTop = 0;
}

async function updateTask(id,task,checked) {
    const resp = await fetch("http://localhost:3001/updatetask",{
        method:"PUT",
        headers:{
            'Content-Type':'application/json'
        },
        body: JSON.stringify({id: id,task:task,status:checked})
    });

    const data = await resp.json()
    if (data.message=="success"){
        await loadTask();
    }
    else{
        alert(data.message)
    }
}
async function deleteallTask(){

    const resp = await fetch("http://localhost:3001/deletetask",{
        method:"DELETE",
        headers:{
            'Content-Type':'application/json'
        },
        
    });

    const data = await resp.json()

    if (data.message=="success"){
        console.log("data deleted");
    }
    else{
        alert(data.message);
    }
    await loadTask();
}