require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
require( 'datatables.net-bs4' )( window, $ );

// Set using https://datatables.net/

$(document).ready( function () {
    $('#eventTable').DataTable( {
        "columnDefs": [
            { "orderable": false, "targets": 2 }
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
} );