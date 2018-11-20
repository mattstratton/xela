require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
// * Set using https://datatables.net/
var dt = require( 'datatables.net-bs4' )( window, $ );
// * Set using https://flatpickr.js.org/
var flatpickr = require("flatpickr");
require("../../node_modules/flatpickr/dist/flatpickr.min.css");


$(document).ready( function () {
    flatpickr('#datepicker', {
        allowInput: true,
        defaultDate: "today",
    });
    $('#eventTable').DataTable( {
        "columnDefs": [
            { "orderable": false, "targets": 3 }
          ]
    });
    $('#abstractTable').DataTable( {
        "columnDefs": [
            { "orderable": false, "targets": 2 }
          ]
    });
    $('#dutonianTable').DataTable( {
        "columnDefs": [
            { "orderable": false, "targets": 1 }
          ]
    });
    $('#proposalTable').DataTable( {
        "columnDefs": [
            { "orderable": false, "targets": 4 }
          ]
    });
    $('#sponsorshipTable').DataTable( {
        "columnDefs": [
            { "orderable": false, "targets": 3 }
          ]
    });
    $('#linksTable').DataTable( {
        "columnDefs": [
            { "orderable": false, "targets": 3 }
          ]
    });

} );