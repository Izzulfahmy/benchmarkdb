<?php

use Illuminate\Support\Facades\Route;
use Illuminate\Support\Facades\DB;

Route::get('/users', function () {
    return DB::table('benchmark_users')->select('id', 'name', 'email')->limit(10)->get();
});
