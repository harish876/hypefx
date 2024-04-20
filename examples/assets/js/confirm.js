/**
 * @typedef {"deleteRow" | "updateRow" | "saveRow"} ConfirmActions
 */
document.addEventListener("htmx:confirm",
    /**
     * @param {{target:{id:ConfirmActions}}} e 
     */
    function (e) {
        const id = e?.target?.id
        if (id === "deleteRow") {
            console.log("Triggered", e)
            e.preventDefault()
            Swal.fire({
                title: "Are you sure?",
                text: "You won't be able to revert this!",
                icon: "warning",
                showCancelButton: true,
                confirmButtonColor: "#3085d6",
                cancelButtonColor: "#d33",
                confirmButtonText: "Yes, delete it!"
            }).then(
                function (result) {
                    if (result.isConfirmed) {
                        Swal.fire({
                            title: "Deleted!",
                            text: "Your file has been deleted.",
                            icon: "success"
                        });
                        e.detail.issueRequest(true)
                    }
                })
        }
        else if (id === "updateRow") {
            let editing = document.querySelector('.editing')
            if (editing) {
                Swal.fire({
                    title: 'Already Editing',
                    showCancelButton: true,
                    confirmButtonText: 'Yep, Edit This Row!',
                    text: 'Hey!  You are already editing a row!  Do you want to cancel that edit and continue?'
                })
                    .then((result) => {
                        if (result.isConfirmed) {
                            htmx.trigger(editing, 'cancel')
                            htmx.trigger(this, 'edit')
                        }
                    })
            } else {
                htmx.trigger(this, 'edit')
            }
        }
        else if (id === "saveRow") {
            console.log("Save Row")
            e.preventDefault()
            Swal.fire({
                title: "Are you sure?",
                text: "You won't be able to revert this!",
                icon: "warning",
                showCancelButton: true,
                confirmButtonColor: "#3085d6",
                cancelButtonColor: "#d33",
                confirmButtonText: "Yes, Save it!"
            }).then((result) => {
                if (result.isConfirmed) {
                    Swal.fire({
                        title: "Saved!",
                        text: "Row has been edited Successfully.",
                        icon: "success"
                    });
                    e.detail.issueRequest(true)
                }
            });
        }
    })


    