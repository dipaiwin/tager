$(document).ready(() => {
    let locCity = null;
    let isMostPop = true;
    $(document).on("click","#isAvgPop", function () {
        isMostPop=false;
    });
    $(document).on("click","#isMostPop", function () {
        isMostPop=true;
    });
    $(document).on("click", ".browse", function () {
        var file = $(this).parents().find(".file");
        file.trigger("click");
    });
    $('input[type="file"]').change(function (e) {
        var fileName = e.target.files[0].name;
        $("#file").val(fileName);

        var reader = new FileReader();
        reader.onload = function (e) {
            // get loaded data and render thumbnail.
            document.getElementById("preview").src = e.target.result;
        };
        // read the image file as a data URL.
        reader.readAsDataURL(this.files[0]);
    });
    $(".copyBtn").click(function () {
        $(".tagsLine").select();
        document.execCommand('copy');
    });
    $('.mainForm').submit(function (event) {
        $(".sendBtn").attr("disabled", "disabled");
        $(".sendBtn").html("<span class=\"spinner-grow spinner-grow-sm\" role=\"status\" aria-hidden=\"true\"></span>");
        $(".sendBtn").append(" Loading...");
        event.preventDefault();
        let formData = new FormData($(this)[0]);
        formData.append("location", locCity);
        formData.append("isMostPop", isMostPop);
        console.log(formData);
        $.ajax({
            url: '/client/predict',
            type: 'POST',
            data: formData,
            cache: false,
            contentType: false,
            processData: false,
            success: function (result) {
                let objResult = JSON.parse(result);
                let textOut = "";
                for (let val of objResult) {
                    textOut += "#" + val + " ";
                }
                console.log(textOut);
                let tagLine = $('.tagsLine');
                tagLine.val(textOut);
                let newHeight = (20 + tagLine[0].scrollHeight) + 'px';
                $('.tagsLine').attr('style', 'height: ' + newHeight);
                $(".sendBtn").removeAttr("disabled");
                $(".sendBtn").text("Отправить");
            },
            error: function (error) {
                alert(error.responseText);
                $(".sendBtn").removeAttr("disabled");
                $(".sendBtn").text("Отправить");
            }
        });
    });
});