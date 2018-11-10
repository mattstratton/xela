require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
require( 'datatables.net-bs4' )( window, $ );

// Set using https://datatables.net/

$(document).ready( function () {
    console.log($);
    $('#eventTable').DataTable( {
        "columnDefs": [
            { "orderable": false, "targets": 2 }
          ]
    });
} );