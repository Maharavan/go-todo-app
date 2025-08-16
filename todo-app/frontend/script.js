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
        if (val==''){
            throw new Error('Empty string');
        }
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
        const sec = document.createElement('div');
        sec.classList.add('check-delete-align');
        const label = document.createElement("label");
        label.htmlFor = taskout.id;
        label.textContent = taskout.task;

        const out = document.createElement("input");
        out.type = 'checkbox';
        out.id = taskout.id;
        out.name = "task-align"
        out.value = taskout.task
        out.checked = taskout.status;
        if (out.checked) {
            main.style.background = "linear-gradient(135deg, #6a11cb, #2575fc)"; // Purple-Blue
        } else {
            main.style.background = "linear-gradient(135deg, #ad60ffff, #93c5fd)"; // Soft Pink-Peach
        }

        // Update background whenever checkbox changes
        out.addEventListener("change", () => {
            if (out.checked) {
                main.style.background = "linear-gradient(135deg, #6a11cb, #2575fc)";
            } else {
                main.style.background = "linear-gradient(135deg, #ad60ffff, #93c5fd)";
            }
            updateTask(taskout.id, taskout.task, out.checked);
        });

        const delete_task = document.createElement("button")
        delete_task.classList.add('task-delete')
        delete_task.innerHTML = `<img src="assets/delete.png" alt="delete">`;
        delete_task.addEventListener("click",()=>{
            delete_required_task(taskout.id) 
        });
        main.appendChild(label);

        sec.appendChild(out);
        sec.appendChild(delete_task);
        main.appendChild(sec);

        taskList.appendChild(main);
    });
    taskList.scrollTop = 0;
}

async function delete_required_task(id) {
    let proceed = confirm('Do you want to delete?')
    if (proceed){
        const resp = await fetch(`http://localhost:3001/deletetask?id=${id}`,{
            method:"DELETE",
            headers:{
                'Content-Type':'application/json'
            }
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
    let proceed = confirm('Do you want to clear all task?');
    if (proceed){
        const resp = await fetch("http://localhost:3001/clearalltask",{
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
}