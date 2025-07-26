package usecase

import "fmt"

func GetCorreoInscripcionRealizada(nombre, modalidad, diasAsistencia string) string {

	if modalidad == "presencial" {
		if diasAsistencia == "sabado" {
			return correoSabadoPresencial(nombre)
		}

		return correoViernesDomingo(nombre)
	}

	return correoSabadoVirtual(nombre)
}

func correoSabadoPresencial(nombre string) string {
	return fmt.Sprintf(`
		<p style="font-family: Arial, sans-serif; font-size: 14px; color: #333;">
			Estimado(a) <strong>%s</strong>,
		</p>

		<p style="font-family: Arial, sans-serif; font-size: 14px; color: #333;">
			Nos complace informarle que su inscripción al evento <strong>Aniversario #25 de la Iglesia La Gran Comisión – Comunidad Cristiana Integral</strong> ha sido registrada exitosamente.
		</p>

		<p style="font-family: Arial, sans-serif; font-size: 14px; color: #333;">
			<strong>Modalidad de inscripción:</strong> Presencial<br>
			<strong>Día de asistencia:</strong> Sábado
		</p>

		<p style="font-family: Arial, sans-serif; font-size: 14px; color: #333;">
			El seminario aplica para personas mayores de 12 años. Esta jornada incluirá los siguientes espacios:
		</p>

		<ul style="font-family: Arial, sans-serif; font-size: 14px; color: #333; padding-left: 20px;">
			<li>Ponencias</li>
			<li>Panel con sesión de preguntas y respuestas</li>
			<li>Merienda</li>
			<li>Almuerzo</li>
			<li>Obsequio especial</li>
		</ul>

		<p style="font-family: Arial, sans-serif; font-size: 14px; color: #333;">
			Para mantenerse informado(a) sobre actualizaciones y detalles logísticos del evento, le invitamos a unirse al grupo oficial de WhatsApp a través del siguiente enlace:<br>
			<a href="https://chat.whatsapp.com/KjwFoapjkUY73hL8abjJ4X" style="color: #00349a;">Unirse al grupo de WhatsApp</a>
		</p>

		<p style="font-family: Arial, sans-serif; font-size: 14px; color: #333;">
			Adjunto encontrará un archivo PDF con la programación completa del evento.
		</p>

		<p style="font-family: Arial, sans-serif; font-size: 14px; color: #333;">
			Si requiere mayor información, no dude en comunicarse con nosotros a través de los siguientes medios:
		</p>

		<ul style="font-family: Arial, sans-serif; font-size: 14px; color: #333; padding-left: 20px;">
			<li>Correo electrónico: <a href="mailto:grancomisionccieventos@gmail.com" style="color: #00349a;">grancomisionccieventos@gmail.com</a></li>
			<li>WhatsApp: <a href="https://wa.me/573166972613" style="color: #00349a;">316 697 2613</a></li>
		</ul>

		<p style="font-family: Arial, sans-serif; font-size: 14px; color: #333;">
			Agradecemos su participación y esperamos compartir juntos este tiempo especial.
		</p>

		<p style="font-family: Arial, sans-serif; font-size: 14px; color: #00349a;">
			Atentamente,<br>
			Iglesia La Gran Comisión – CCI</strong>
		</p>
	`, nombre)
}

func correoSabadoVirtual(nombre string) string {
	return fmt.Sprintf(`
		<p style="font-family: Arial, sans-serif; font-size: 14px; color: #333;">
			Estimado(a) <strong>%s</strong>,
		</p>

		<p style="font-family: Arial, sans-serif; font-size: 14px; color: #333;">
			Le confirmamos que su inscripción al evento <strong>Aniversario #25 de la Iglesia La Gran Comisión – Comunidad Cristiana Integral</strong> ha sido registrada exitosamente.
		</p>

		<p style="font-family: Arial, sans-serif; font-size: 14px; color: #333;">
			<strong>Modalidad de inscripción:</strong> Virtual<br>
			<strong>Día de asistencia:</strong> Sábado
		</p>

		<p style="font-family: Arial, sans-serif; font-size: 14px; color: #333;">
			Para mantenerse informado(a) sobre las instrucciones de acceso y novedades del evento, le invitamos a unirse al grupo oficial de WhatsApp a través del siguiente enlace:<br>
			<a href="https://chat.whatsapp.com/KjwFoapjkUY73hL8abjJ4X" style="color: #00349a;">Unirse al grupo de WhatsApp</a>
		</p>

		<p style="font-family: Arial, sans-serif; font-size: 14px; color: #333;">
			Adjunto encontrará un archivo PDF con la programación detallada del evento.
		</p>

		<p style="font-family: Arial, sans-serif; font-size: 14px; color: #333;">
			Si requiere mayor información o tiene inquietudes, puede comunicarse con nosotros a través de los siguientes medios:
		</p>

		<ul style="font-family: Arial, sans-serif; font-size: 14px; color: #333; padding-left: 20px;">
			<li>Correo electrónico: <a href="mailto:grancomisionccieventos@gmail.com" style="color: #00349a;">grancomisionccieventos@gmail.com</a></li>
			<li>WhatsApp: <a href="https://wa.me/573166972613" style="color: #00349a;">316 697 2613</a></li>
		</ul>

		<p style="font-family: Arial, sans-serif; font-size: 14px; color: #333;">
			Agradecemos su participación y esperamos contar con su asistencia virtual.
		</p>

		<p style="font-family: Arial, sans-serif; font-size: 14px; color: #00349a;">
			Atentamente,<br>
			Iglesia La Gran Comisión – CCI</strong>
		</p>
	`, nombre)
}

func correoViernesDomingo(nombre string) string {
	return fmt.Sprintf(`
		<p style="font-family: Arial, sans-serif; font-size: 14px; color: #333;">
			Estimado(a) <strong>%s</strong>,
		</p>

		<p style="font-family: Arial, sans-serif; font-size: 14px; color: #333;">
			Nos complace informarle que su inscripción al evento <strong>Aniversario #25 de la Iglesia La Gran Comisión – Comunidad Cristiana Integral</strong> ha sido registrada exitosamente.
		</p>

		<p style="font-family: Arial, sans-serif; font-size: 14px; color: #333;">
			<strong>Modalidad de inscripción:</strong> Presencial<br>
			<strong>Días de asistencia:</strong> Viernes y/o Domingo
		</p>

		<p style="font-family: Arial, sans-serif; font-size: 14px; color: #333;">
			Para recibir información actualizada y detalles logísticos del evento, le invitamos a unirse al grupo oficial de WhatsApp a través del siguiente enlace:<br>
			<a href="https://chat.whatsapp.com/KjwFoapjkUY73hL8abjJ4X" style="color: #00349a;">Unirse al grupo de WhatsApp</a>
		</p>

		<p style="font-family: Arial, sans-serif; font-size: 14px; color: #333;">
			Adjunto a este mensaje encontrará un archivo PDF con la programación completa del evento.
		</p>

		<p style="font-family: Arial, sans-serif; font-size: 14px; color: #333;">
			Para cualquier inquietud o información adicional, puede comunicarse con nosotros a través de los siguientes medios:
		</p>

		<ul style="font-family: Arial, sans-serif; font-size: 14px; color: #333; padding-left: 20px;">
			<li>Correo electrónico: <a href="mailto:grancomisionccieventos@gmail.com" style="color: #00349a;">grancomisionccieventos@gmail.com</a></li>
			<li>WhatsApp: <a href="https://wa.me/573166972613" style="color: #00349a;">316 697 2613</a></li>
		</ul>

		<p style="font-family: Arial, sans-serif; font-size: 14px; color: #333;">
			Agradecemos su participación. Será un honor contar con su presencia durante este tiempo de celebración y edificación.
		</p>

		<p style="font-family: Arial, sans-serif; font-size: 14px; color: #00349a;">
			Atentamente,<br>
			Iglesia La Gran Comisión – CCI</strong>
		</p>
	`, nombre)
}
