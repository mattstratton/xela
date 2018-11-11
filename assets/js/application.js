require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
// * Set using https://datatables.net/
var dt = require( 'datatables.net-bs4' )( window, $ );
var dp = require('bootstrap-datepicker');


$(document).ready( function () {
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
    // https://bootstrap-datepicker.readthedocs.io/en/stable/index.html
    $('.datepicker').datepicker()
} );