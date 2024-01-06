create or replace function cancelar_turnos(dni_medique_pedide int, fecha_desde date, fecha_hasta date) returns int as $$
declare
	cont_turnos_cancelados int;
	nro_turno_cancelado int;
	v record;
begin 
	cont_turnos_cancelados = 0;
	for v in select * from turno where turno.dni_medique = dni_medique_pedide and (estado = 'disponible' or estado = 'reservado') loop
		if v.fecha >= fecha_desde and v.fecha <= fecha_hasta + '1 day'::interval then

			update turno set estado = 'cancelado' where v.nro_turno = turno.nro_turno;
	
			select nro_turno into nro_turno_cancelado from turno where nro_turno = v.nro_turno;
			
			perform reprogramar(nro_turno_cancelado);

			cont_turnos_cancelados = cont_turnos_cancelados + 1;

		end if;
	end loop;	
	return cont_turnos_cancelados;
end;
$$ language plpgsql;

create or replace function reprogramar(nro_turno_cancelado int) returns void as $$
declare
	paciente_selec record;
	turno_cancelado record;
	medique_selec record;
begin
	select * into turno_cancelado from turno where turno.nro_turno = nro_turno_cancelado;
	select * into medique_selec from medique where dni_medique = turno_cancelado.dni_medique;
	select * into paciente_selec from paciente where nro_paciente = turno_cancelado.nro_paciente;
		
	insert into reprogramacion values (turno_cancelado.nro_turno, paciente_selec.nombre, 
		paciente_selec.apellido, paciente_selec.telefono, paciente_selec.email, medique_selec.nombre, medique_selec.apellido, 'pendiente');
	end;
$$ language plpgsql;
