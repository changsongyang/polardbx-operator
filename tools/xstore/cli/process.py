# Copyright 2021 Alibaba Group Holding Limited.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import os
import signal
import sys
import subprocess

import click

from .common import global_mgr


@click.group(name='process')
def process_group():
    pass


def _send_signal_to_process(pid: int or None, sig: int):
    if not pid:
        return
    print('kill process %d with signal %d', pid, sig)
    os.kill(pid, sig)


@click.command(name='kill')
def kill_engine_process():
    _send_signal_to_process(global_mgr.engine().engine_process_id(), signal.SIGKILL)


process_group.add_command(kill_engine_process)


@click.command(name='stop')
def stop_engine_process():
    _send_signal_to_process(global_mgr.engine().engine_process_id(), signal.SIGSTOP)


process_group.add_command(stop_engine_process)


@click.command(name='continue')
def continue_engine_process():
    _send_signal_to_process(global_mgr.engine().engine_process_id(), signal.SIGCONT)


process_group.add_command(continue_engine_process)


@click.command(name='kill_all_mysql')
def kill_engine_all_process():
    subprocess.Popen(
        ["/usr/bin/sh", "-c",
         "pid=`ps -ef|grep /tools/xstore/current/entrypoint | grep -v grep | awk '{print $2}'`&&kill -9 $pid"],
        cwd=None, stdout=None, stderr=None)


process_group.add_command(kill_engine_all_process)


@click.command(name='kill_mysqld')
def kill_process_mysqld():
    subprocess.Popen(
        ["/usr/bin/sh", "-c",
         "pid=`ps -ef |grep mysqld | grep  loose-pod-name |  grep defaults-file | grep -v mysqld_safe | awk '{print "
         "$2}'`&&kill $pid"],
        cwd=None, stdout=None, stderr=None)


process_group.add_command(kill_process_mysqld)


@click.command(name='check_std_err_complete')
@click.option('--filepath', type=str)
def check_std_err_complete(filepath: str, keyword="completed OK!"):
    complete = False
    with open(filepath, 'r') as f:
        for line in f.readlines():
            if line.strip().endswith(keyword):
                complete = True
    if complete:
        sys.exit(0)
    sys.exit(1)


process_group.add_command(check_std_err_complete)
