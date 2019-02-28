$(function () {

    var data = {
        "category": ""
    }

    var dataTableLang = {
        "processing": "Подождите...",
        "search": "Поиск:",
        "lengthMenu": "Показать _MENU_ записей",
        "info": "Записи с _START_ до _END_ из _TOTAL_ записей",
        "infoEmpty": "Записи с 0 до 0 из 0 записей",
        "infoFiltered": "(отфильтровано из _MAX_ записей)",
        "infoPostFix": "",
        "loadingRecords": "Загрузка записей...",
        "zeroRecords": "Записи отсутствуют.",
        "emptyTable": "В таблице отсутствуют данные",
        "paginate": {
            "first": "Первая",
            "previous": "Предыдущая",
            "next": "Следующая",
            "last": "Последняя"
        },
        "aria": {
            "sortAscending": ": активировать для сортировки столбца по возрастанию",
            "sortDescending": ": активировать для сортировки столбца по убыванию"
        }
    }

    var dataTableCfg = {
        "destroy": true,
        "language": dataTableLang,
        "processing": true,
        "serverSide": false,
        "ajax": {
            "url": "api/v1/dynamic_content/",
            "type": "POST",
            "data": getData
        },
        "columns": [
            { "data": "id" },
            { "data": "date" },
            { "data": "title" },
            { "data": "category" },
        ]
    }

    var table = $("#dynamic-table").DataTable(dataTableCfg);
    function getData() {
        return data
    }

    var datepicker = $("#date-range").datepicker().data('datepicker');
    function setSortParams() {
        var startDate;
        var endDate;
        if (datepicker.selectedDates[0] != null) startDate = new Date(datepicker.selectedDates[0]).toISOString();
        if (datepicker.selectedDates[1] != null) endDate = new Date(datepicker.selectedDates[1]).toISOString();
        data = {
            "category": $('#select-category').val(),
            "startDate": startDate,
            "endDate": endDate
        }
        table.ajax.reload()
    }
    function clearSortParams() {
        datepicker.clear()
        $('#select-category').val("")
    }
    function exportData() {
        openCmsMainCatDir = $('input#output-main-category').val();
        openCmsTargetCatDir = $('input#output-category').val();
        openCmsOutputFormat = $('select#select-opencms-type').val();
        $.post("/api/v1/convert_dynamic_content/", {
            "output-main-category": openCmsMainCatDir,
            "output-category": openCmsTargetCatDir,
            "select-opencms-type": openCmsOutputFormat
        });
    }

    $('#apply-sort-params').click(setSortParams);
    $('#clear-sort-params').click(clearSortParams);
    $('#export-btn').click(exportData);

})